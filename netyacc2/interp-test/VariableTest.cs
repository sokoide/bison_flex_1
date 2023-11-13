using System.Runtime.Intrinsics;
using interp_lib.Interp;

namespace interp_test;

public class VariableTest
{
    [Fact]
    public void Variable_Basic()
    {
        Variable v1 = new Variable(VariableType.INT, "hoge");
        Variable v2 = new Variable(VariableType.INT, "page");
        Assert.NotEqual(v1, v2);

        v2 = new Variable(VariableType.STRING, "hoge");
        Assert.Equal(v1, v2);
    }

    [Fact]
    public void Variable_Dictionary()
    {
        Dictionary<Variable, int> d = new Dictionary<Variable, int>();

        Variable v1 = new Variable(VariableType.INT, "hoge");
        Variable v2 = new Variable(VariableType.INT, "page");
        d[v1] = 42;
        d[v2] = 43;

        Variable v3 = new Variable(VariableType.STRING, "hoge");
        Assert.Equal(42, d[v3]);
    }
}
