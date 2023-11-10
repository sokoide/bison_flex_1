namespace netyacc2.Interp
{
    class Program
    {
        static void Main(string[] args)
        {
            var parser = new Interp.InterpParser();
            var input = @"/*a=-123;
put(a);
a=40+2;
put(a);
a=40-2;
put(a);
b=40*2;
put(b);
b=40/2;
put(b);
*/
i=3;
while(i>0){
    put(i);
    i=i-1;
}
/*
i=2;
while(i){
    put(i);
    i=i-1;
}*/
";
            Console.WriteLine("* input");
            Console.WriteLine(input);
            Console.WriteLine("* Parse");
            parser.Parse(input);
            Console.WriteLine("* Execute");
            parser.Execute();
        }
    }
}