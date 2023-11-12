using System.Text;
using interp_lib.Interp;

namespace interp_exe;

public class Exe
{
    public static int Main(string[] args)
    {
        int ret = 0;
        var parser = new InterpParser();
        var vm = new VM();

        if (args.Length > 0 && args[0] == "demo")
        {
            ret = Demo(parser, vm);
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
            ret = vm.Execute(resolvedCode, parser.ItoS);
        }

        return ret;
    }

    public static int Demo(InterpParser parser, VM vm)
    {
        var input = @"put(""*** Demo ***"");
put(""counting down..."");
e = 3;
while (e > 0)
{
    put(""e="", e);
    e = e - 1;
    }
";
        parser.Parse(input);

        var resolvedCode = vm.ResoleLabels(parser.Code);
        Console.WriteLine("* Source");
        Console.WriteLine(input);
        Console.WriteLine("* Original. Jump/JumpF's operands mean Label name");
        vm.Dump(parser.Code);
        Console.WriteLine("* Label Resolved. Jump/JumpF's operands mean PC");
        vm.Dump(resolvedCode);
        Console.WriteLine("* String table");
        vm.DumpStringTable(parser.ItoS);
        Console.WriteLine("* Executing...");
        return vm.Execute(resolvedCode, parser.ItoS);
    }
}