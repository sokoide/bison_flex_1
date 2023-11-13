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
    [InlineData("int a; int b; string c; a=42; return a;", 42)]
    [InlineData("int a; int b; a=42; return b;", 0)]
    [InlineData("x=123; x=x+1; return x;", 124)]
    [InlineData("y=456; x=2; y=(y+4)/x; return y;", 230)]
    public void VM_Execute_Basic(string input, int want)
    {
        parser.Parse(input);
        var resolvedCode = vm.ResoleLabels(parser.Code);
        int got = vm.Execute(resolvedCode, parser.ItoS, parser.ItoV);
        Assert.Equal(want, got);
    }

    [Theory]
    [InlineData("a=1; b=2; x=3; if(x>=3) { return 1; } return 0;", 1)]
    [InlineData("x=0; if(x>0) {return 1;} else {return 2;}", 2)]
    public void VM_Execute_If(string input, int want)
    {
        parser.Parse(input);
        var resolvedCode = vm.ResoleLabels(parser.Code);
        int got = vm.Execute(resolvedCode, parser.ItoS, parser.ItoV);
        Assert.Equal(want, got);
    }

    [Theory]
    [InlineData("a=3; b=100; while(a>0){b=b+1; a=a-1;} return b;", 103)]
    public void VM_Execute_While(string input, int want)
    {
        parser.Parse(input);
        var resolvedCode = vm.ResoleLabels(parser.Code);
        int got = vm.Execute(resolvedCode, parser.ItoS, parser.ItoV);
        Assert.Equal(want, got);
    }
}
