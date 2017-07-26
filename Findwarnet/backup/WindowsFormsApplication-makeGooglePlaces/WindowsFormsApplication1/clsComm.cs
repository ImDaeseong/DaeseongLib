using System.Windows.Forms;

namespace WindowsFormsApplication1
{
    class clsComm
    {
        private static clsComm selfInstance = null;
        public static clsComm getInstance
        {
            get
            {
                if (selfInstance == null) selfInstance = new clsComm();
                return selfInstance;
            }
        }

        public Form1 GetMainForm()
        {
            foreach (Form fFind in Application.OpenForms)
            {
                if (fFind is Form1)
                    return fFind as Form1;
            }
            return null;
        }

        public string NullVal(object src, string Value)
        {
            if (src != null)
                return src.ToString();
            return Value;
        }

        public string QStr(string sValue)
        {
            return "'" + sValue.Replace("'", "''") + "'";
        }

    }
}
