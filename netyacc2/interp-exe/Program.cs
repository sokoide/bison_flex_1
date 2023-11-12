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
        var input = "e=3;\nwhile(e>0)\n{\n\tput(e);\n\te=e-1;\n}";
        parser.Parse(input);

        var resolvedCode = vm.ResoleLabels(parser.Code);
        Console.WriteLine("* Source");
        Console.WriteLine(input);
        Console.WriteLine("* Original. Jump/JumpF's operands mean Label name");
        vm.Dump(parser.Code);
        Console.WriteLine("* Label Resolved. Jump/JumpF's operands mean PC");
        vm.Dump(resolvedCode);
        Console.WriteLine("* Executing...");
        vm.Execute(resolvedCode);
    }
}