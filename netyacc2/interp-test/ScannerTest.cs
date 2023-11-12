using Xunit.Abstractions;
using interp_lib.Interp;

namespace interp_test;
public class ScannerTest : IDisposable
{
    internal InterpScanner tgt;
    private readonly ITestOutputHelper output;

    public ScannerTest(ITestOutputHelper output)
    {
        this.tgt = new InterpScanner();
        this.output = output;
    }

    public void Dispose()
    {
    }


    [Fact]
    public void AssemblyName()
    {
        Assert.Equal("interp_test", System.Reflection.Assembly.GetExecutingAssembly().GetName().Name);
    }

    [Fact]
    public void Scanner_Pool()
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