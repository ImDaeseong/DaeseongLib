using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using System.Runtime.InteropServices;

namespace MouseHook
{
    public partial class Form1 : Form
    {
        private struct POINT
        {
            public int x;
            public int y;
        }

        [DllImport("user32.dll", SetLastError = true, CharSet = CharSet.Auto)]
        private static extern IntPtr WindowFromPoint(POINT lpPoint);

        [DllImport("user32.dll", SetLastError = true, CharSet = CharSet.Auto)]
        private static extern int GetClassName(IntPtr hWnd, StringBuilder lpClassName, int nMaxCount);

        [DllImport("user32.dll", SetLastError = true, CharSet = CharSet.Auto)]
        private static extern int GetWindowText(IntPtr hWnd, StringBuilder lpString, int nMaxCount);

        
        private const int WM_USER = 0x0400;
        private const int PT_GETWINDOWFROMPOINT = WM_USER + 1;


        protected override void WndProc(ref Message m)
        {
            switch (m.Msg)
            {
                case PT_GETWINDOWFROMPOINT:
                    {
                        POINT pt;
                        pt.x = m.WParam.ToInt32();
                        pt.y = m.LParam.ToInt32();

                        IntPtr hwnd = WindowFromPoint(pt);
                        if (hwnd != IntPtr.Zero)
                        {
                            StringBuilder title = new StringBuilder(256);
                            GetWindowText(hwnd, title, title.Capacity);

                            StringBuilder className = new StringBuilder(256);
                            GetClassName(hwnd, className, className.Capacity);

                            Console.WriteLine("title:" + title);
                            Console.WriteLine("className:" + className);
                        }                        
                    }
                    break;
            }
            base.WndProc(ref m);
        }

        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
           
        }

        private void Form1_FormClosed(object sender, FormClosedEventArgs e)
        {
            
        }
    }

   
}
