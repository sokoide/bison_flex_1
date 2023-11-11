using interp_lib.Interp;

namespace interp_exe;

public class Exe
{
    public static void Main(string[] args)
    {
        var parser = new interp_lib.Interp.InterpParser();
        var vm = new interp_lib.Interp.VM();

        // var input = "a=42;put(a);";
        var input = "a=3;while(a>0){put(a);a=a-1;}";
        parser.Parse(input);

        var resolvedCode = vm.ResoleLabels(parser.Code);
        Console.WriteLine("* Original");
        vm.Dump(parser.Code);
        Console.WriteLine("* Label Resolved");
        vm.Dump(resolvedCode);
        Console.WriteLine("* Executing...");
        vm.Execute(resolvedCode);
    }
}