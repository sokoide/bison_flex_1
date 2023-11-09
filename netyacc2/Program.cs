namespace netyacc2.Interp
{
    class Program
    {
        static void Main(string[] args)
        {
            var parser = new Interp.InterpParser();
            var input = @"i=-123;
print(40+2);
print(40-2);
print(40*2);
/* comment */
print(40/2);
";
            Console.WriteLine(input);
            parser.Parse(input);
        }
    }
}