using System.Reflection.Metadata.Ecma335;

namespace interp_lib.Interp
{
    public enum VariableType
    {
        INT,
        STRING,
    }

    public class Variable
    {
        public VariableType Vt;
        public string Name = "";

        public Variable(VariableType vt, string name)
        {
            Vt = vt;
            Name = name;
        }

        public override string ToString()
        {
            return $"Name: {Name}, Type: {Vt}";
        }

        public override int GetHashCode()
        {
            return Name.GetHashCode();
        }

        public override bool Equals(object? obj)
        {
            var other = obj as Variable;
            if (other == null) return false;

            return Name == other.Name;
        }
    }
}