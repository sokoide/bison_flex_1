using interp_lib;
using Microsoft.VisualStudio.TestPlatform.Utilities;
using Xunit.Abstractions;
using interp_lib.Interp;

namespace interp_test;

public class ParserTest : IDisposable
{
    private interp_lib.Interp.InterpParser t;
    private readonly ITestOutputHelper output;

    public ParserTest(ITestOutputHelper output)
    {
        this.t = new interp_lib.Interp.InterpParser();
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
        t.Parse(input);
    }

    [Fact]
    public void Parser_Exception()
    {
        string input = @"a if then;";

        var exc = Assert.Throws<Exception>(() =>
        t.Parse(input));

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
        t.Parse(input));

        string got = exc.Message.Substring(0, want.Length);
        // test if got starts with want
        Assert.Equal(want, got);
    }

    [Fact]
    public void Parser_GeneratedCode()
    {
        string input = @"foo=42;";
        t.Parse(input);
        Assert.Equal(2, t.Code.Count);
        Assert.Equal(Op.PushN, t.Code[0].Op);
        Assert.Equal(42, t.Code[0].Sub);
        Assert.Equal(Op.Pop, t.Code[1].Op);
        Assert.Equal('f' - 'a', t.Code[1].Sub);
    }
}