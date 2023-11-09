using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace netyacc1
{
    class Program
    {
        static void Main(string[] args)
        {
            var parser = new Calc.CalcParser();

            do
            {
                Console.Write("> ");
                var input = Console.ReadLine();
                if (string.IsNullOrEmpty(input))
                {
                    return;
                }

                parser.Parse(input);
            } while (true);
        }
    }
}

