using interp_lib;
using Microsoft.VisualStudio.TestPlatform.Utilities;
using Xunit.Abstractions;

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

        string want = "Syntax error, unexpected IF";
        string got = exc.Message.Substring(0, want.Length);
        // test if got starts with want
        Assert.Equal(want, got);
    }

    [Theory]
    [InlineData("a then", "Syntax error, unexprected IDENT at line 1")]
    public void Parser_Exceptions(string input, string want)
    {
        var exc = Assert.Throws<Exception>(() =>
        t.Parse(input));

        string got = exc.Message.Substring(0, want.Length);
        // test if got starts with want
        Assert.Equal(want, got);
    }

}