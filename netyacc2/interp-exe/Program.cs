using System.Text;
using interp_lib.Interp;

namespace interp_exe;

public class Exe
{
    public static void Main(string[] args)
    {
        var parser = new InterpParser();
        var vm = new VM();

        if (args.Length > 0 && args[0] == "demo")
        {
            Demo(parser, vm);
        }
        else
        {
            string line;
            StringBuilder sb = new StringBuilder();
            while ((line = Console.ReadLine()) != null && line != "")
            {
                sb.AppendLine(line);
            }

            parser.Parse(sb.ToString());
            var resolvedCode = vm.ResoleLabels(parser.Code);
            vm.Execute(resolvedCode);
        }
    }

    public static void Demo(InterpParser parser, VM vm)
    {
        var input = "a=3;\nwhile(a>0)\n{\n\tput(a);\n\ta=a-1;\n}";
        parser.Parse(input);

        var resolvedCode = vm.ResoleLabels(parser.Code);
        Console.WriteLine("* Source");
        Console.WriteLine(input);
        Console.WriteLine("* Original");
        vm.Dump(parser.Code);
        Console.WriteLine("* Label Resolved");
        vm.Dump(resolvedCode);
        Console.WriteLine("* Executing...");
        vm.Execute(resolvedCode);
    }
}