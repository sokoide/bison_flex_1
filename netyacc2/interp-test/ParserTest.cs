using Xunit.Abstractions;
using interp_lib.Interp;

namespace interp_test;

public class ParserTest : IDisposable
{
    private InterpParser tgt;
    private readonly ITestOutputHelper output;

    public ParserTest(ITestOutputHelper output)
    {
        this.tgt = new InterpParser();
        this.output = output;
    }

    public void Dispose()
    {
    }

    [Theory]
    [InlineData("a=42;")]
    [InlineData("a=42;b=100;")]
    [InlineData("a=1;if(a==1){put(a);}")]
    public void Parser_BasicSyntax(string input)
    {
        output.WriteLine("* Testing {0}...", input);
        tgt.Parse(input);
    }

    [Fact]
    public void Parser_Exception()
    {
        string input = @"a if then;";

        var exc = Assert.Throws<Exception>(() =>
        tgt.Parse(input));

        // "1:2" means line 1, col 2
        string want = "1:2 Syntax error, unexpected IF";
        string got = exc.Message.Substring(0, want.Length);
        // test if got starts with want
        Assert.Equal(want, got);
    }

    [Theory]
    [InlineData("a then", "1:2 Syntax error, unexpected IDENT")]
    [InlineData("a=1;\nb hoge", "2:2 Syntax error, unexpected IDENT")]
    [InlineData("a=1;\nb=1;\nc hoge", "3:2 Syntax error, unexpected IDENT")]
    public void Parser_Exceptions(string input, string want)
    {
        var exc = Assert.Throws<Exception>(() =>
        tgt.Parse(input));

        string got = exc.Message.Substring(0, want.Length);
        // test if got starts with want
        Assert.Equal(want, got);
    }

    [Fact]
    public void Parser_GeneratedCode()
    {
        string input = @"foo=42;";
        tgt.Parse(input);
        Assert.Equal(2, tgt.Code.Count);
        Assert.Equal(Op.PushN, tgt.Code[0].Op);
        Assert.Equal(42, tgt.Code[0].Sub);
        Assert.Equal(Op.Pop, tgt.Code[1].Op);
        Assert.Equal('f' - 'a', tgt.Code[1].Sub);
    }

    [Theory]
    [InlineData(Op.PushN, Token.NUMBER, 42, 42)]
    [InlineData(Op.Calc, Token.ADD, (int)Token.ADD, (int)Token.ADD)]
    [InlineData(Op.Calc, Token.SUB, (int)Token.SUB, (int)Token.SUB)]
    [InlineData(Op.Calc, Token.MUL, (int)Token.MUL, (int)Token.MUL)]
    [InlineData(Op.Calc, Token.DIV, (int)Token.DIV, (int)Token.DIV)]
    [InlineData(Op.Calc, Token.MINUS, (int)Token.MINUS, (int)Token.MINUS)]
    public void Parser_GenCodeN(Op op, Token token, int n, int wantSub)
    {
        tgt.GenCode(op, new Node(token, n));
        Assert.Single(tgt.Code);
        Assert.Equal(op, tgt.Code[0].Op);
        Assert.Equal(wantSub, tgt.Code[0].Sub);
    }

    [Theory]
    [InlineData(Op.PushI, Token.IDENT, "a", 0)]
    [InlineData(Op.PushI, Token.IDENT, "foo", 5)]
    [InlineData(Op.PushI, Token.IDENT, "zoo", 25)]
    [InlineData(Op.Pop, Token.IDENT, "b", 1)]
    [InlineData(Op.PutI, Token.IDENT, "c", 2)]
    public void Parser_GenCodeS(Op op, Token token, string s, int wantSub)
    {
        tgt.GenCode(op, new Node(token, s));
        Assert.Single(tgt.Code);
        Assert.Equal(op, tgt.Code[0].Op);
        Assert.Equal(wantSub, tgt.Code[0].Sub);
    }

    [Fact]
    public void Parser_Pool()
    {
        int got;
        int want;

        want = 1;
        got = tgt.Pool("hoge");
        Assert.Equal(want, got);

        want = 2;
        got = tgt.Pool("page");
        Assert.Equal(want, got);

        want = 3;
        got = tgt.Pool("foo");
        Assert.Equal(want, got);

        want = 1;
        got = tgt.Pool("hoge");
        Assert.Equal(want, got);
    }

}