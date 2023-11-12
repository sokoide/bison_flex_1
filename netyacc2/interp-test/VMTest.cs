using Xunit.Abstractions;
using interp_lib.Interp;

namespace interp_test;

public class VMTest : IDisposable
{
    private InterpParser parser;
    private VM vm;
    private readonly ITestOutputHelper output;

    public VMTest(ITestOutputHelper output)
    {
        this.parser = new InterpParser();
        this.vm = new VM();
        this.output = output;
    }

    public void Dispose()
    {
    }

    [Theory]
    [InlineData("a=42; return 42;", 42)]
    [InlineData("a=42; return a;", 42)]
    [InlineData("x=123; x=x+1; return x;", 124)]
    [InlineData("y=456; x=2; y=(y+4)/x; return y;", 230)]
    public void VM_Execute_Basic(string input, int want)
    {
        parser.Parse(input);
        var resolvedCode = vm.ResoleLabels(parser.Code);
        int got = vm.Execute(resolvedCode, parser.ItoS);
        Assert.Equal(want, got);
    }
}
