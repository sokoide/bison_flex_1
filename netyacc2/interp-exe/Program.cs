using interp_lib.Interp;

namespace interp_exe;

public class Exe
{
    public static void Main(string[] args)
    {
        var parser = new interp_lib.Interp.InterpParser();
        var vm = new interp_lib.Interp.VM();

        var input = "a=42;put(a);";
        parser.Parse(input);
        vm.Execute(parser.Code);
    }
}