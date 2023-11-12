using System.Diagnostics;

namespace interp_lib.Interp
{
    public static class Debug
    {
        public static bool Enabled = false;

        public static void Print(string s)
        {
            if (Enabled)
            {
                var st = new StackTrace();
                var sf = st.GetFrame(1);
                var cn = sf.GetMethod().ReflectedType.Name ?? "";
                var mn = sf.GetMethod().Name ?? "";
                Console.WriteLine("[{0}::{1}] {2}", cn, mn, s);
            }
        }
        public static void Print(string format, params object[] opts)
        {
            if (Enabled)
            {
                var st = new StackTrace();
                var sf = st.GetFrame(1);
                var cn = sf.GetMethod().ReflectedType.Name;
                var mn = sf.GetMethod().Name;
                var s = String.Format(format, opts);
                Console.WriteLine("[{0}::{1}] {2}", cn, mn, s);
            }
        }
    }
}