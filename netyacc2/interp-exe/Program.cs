using System.ComponentModel.Design;
using System.Security.Cryptography;
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
        bool demo = false;
        bool debug = false;
        string inputFilePath = "";
        string input;

        for (int i = 0; i < args.Length; i++)
        {
            string arg = args[i].ToLower();
            if (arg == "--verbose")
            {
                debug = true;
            }
            else if (arg == "demo")
            {
                debug = true;
                demo = true;
            }
            else if (arg == "--file")
            {
                if (i + 1 >= args.Length)
                {
                    Console.WriteLine("please specify a script path");
                    return 1;
                }
                i++;
                inputFilePath = args[i];

            }
            else if (arg == "--help")
            {
                Console.WriteLine("dotnet run [--verbose] [--file scriptpath] [demo]");
                Console.WriteLine(" verbose: verbose output");
                Console.WriteLine(" file:    script file path");
                Console.WriteLine(" demo:    demo mode");
                return 0;
            }
        }

        if (demo)
        {
            input = @"put(""*** Demo ***"");
put(""counting down..."");
e = 3;
while (e > 0)
{
    put(""e="", e);
    e = e - 1;
    }
";
        }
        else
        {
            input = File.ReadAllText(inputFilePath);
        }

        parser.Parse(input);
        var resolvedCode = vm.ResoleLabels(parser.Code);

        if (debug)
        {
            Console.WriteLine("* Source");
            Console.WriteLine(input);
            Console.WriteLine("* Original. Jump/JumpF's operands mean Label name");
            vm.Dump(parser.Code);
            Console.WriteLine("* Label Resolved. Jump/JumpF's operands mean PC");
            vm.Dump(resolvedCode);
            Console.WriteLine("* String table");
            vm.DumpStringTable(parser.ItoS);
            Console.WriteLine();
        }

        ret = vm.Execute(resolvedCode, parser.ItoS);

        return ret;
    }
}