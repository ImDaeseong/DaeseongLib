using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;
using System.IO;
using System.Diagnostics;
using System.Windows.Interop;
using System.Windows.Threading;

namespace Wpfunzip
{
    /// <summary>
    /// MainWindow.xaml에 대한 상호 작용 논리
    /// </summary>
    public partial class MainWindow : Window
    {
        static private DispatcherTimer dispatcherTimer;

        private const int WM_USER = 0x0400;
        private const int ZIP_COMPLETEUNZIP = WM_USER + 1;
        private HwndSource hwndSource;
                
        private string strFileName = "";

        public MainWindow()
        {
            InitializeComponent();

            this.SourceInitialized += MainWindow_SourceInitialized;

            Init();

            InitTimer();
        }

        private void InitTimer()
        {
            dispatcherTimer = new DispatcherTimer();
            dispatcherTimer.Tick += DispatcherTimer_Tick;
            dispatcherTimer.Interval = TimeSpan.FromSeconds(10);
        }
        private void StartTimer()
        {
            dispatcherTimer.Start();
        }

        private void stopTimer()
        {
            dispatcherTimer.Stop();
        }

        private void DispatcherTimer_Tick(object sender, EventArgs e)
        {
            Close();
        }

        private void MainWindow_SourceInitialized(object sender, EventArgs e)
        {
            hwndSource = PresentationSource.FromVisual((Visual)sender) as HwndSource;
            hwndSource.AddHook(new HwndSourceHook(WndProc));
        }

        private IntPtr WndProc(IntPtr hwnd, int msg, IntPtr wParam, IntPtr lParam, ref bool handled)
        {
            if (msg == ZIP_COMPLETEUNZIP)
            {
                Close();
            }
            return IntPtr.Zero;
        }

        string GetFileExtName(string strFilename)
        {
            int nPos = strFilename.LastIndexOf('.');
            int nLength = strFilename.Length;
            if (nPos < nLength)
                return strFilename.Substring(nPos + 1, (nLength - nPos) - 1);
            return string.Empty;
        }

        private void Init(bool bIsEdit = false)
        {            
            if (bIsEdit)
            {
                lblDrag.Visibility = Visibility.Hidden;
                btnUnzip.Visibility = Visibility.Visible;
            }
            else
            {
                lblDrag.Visibility = Visibility.Visible;
                btnUnzip.Visibility = Visibility.Hidden;
            }            
        }

        private void Window_MouseLeftButtonDown(object sender, MouseButtonEventArgs e)
        {
            base.OnMouseLeftButtonDown(e);

            this.DragMove();
        }

        private void Window_Loaded(object sender, RoutedEventArgs e)
        {
            this.AllowDrop = true;
                   
            this.Left = System.Windows.SystemParameters.WorkArea.Left + this.Width - 50;
            this.Top = System.Windows.SystemParameters.WorkArea.Top + this.Height - 50;
        }

        private void Window_Closed(object sender, EventArgs e)
        {
            if (dispatcherTimer != null)
            {
                dispatcherTimer.Stop();
                dispatcherTimer = null;
            }

            Close();
        }

        private void Window_DragOver(object sender, DragEventArgs e)
        {
            if (!e.Data.GetDataPresent(DataFormats.FileDrop))
            {
                e.Effects = DragDropEffects.Move;
                e.Handled = true;
            }
        }

        private void Window_Drop(object sender, DragEventArgs e)
        {
            strFileName = "";
            if (e.Data.GetDataPresent(DataFormats.FileDrop))
            {
                string[] files = (string[])e.Data.GetData(DataFormats.FileDrop);
                foreach (string file in files)
                {
                    if (GetFileExtName(file).ToLower() == "zip")
                    {
                        strFileName = file;
                        break;
                    }
                }
            }
            
            Button_Click(null, null);
        }

        private void Button_Click(object sender, RoutedEventArgs e)
        {
            try
            {
                strFileName = strFileName.Replace(" ", "<br>");

                string strGo = string.Format("{0}Daeseonglib.exe", AppDomain.CurrentDomain.BaseDirectory);
                System.Diagnostics.Process goProcess = new System.Diagnostics.Process();
                goProcess.StartInfo.FileName = strGo;
                goProcess.StartInfo.Arguments = " " + strFileName;
                goProcess.StartInfo.WindowStyle = System.Diagnostics.ProcessWindowStyle.Hidden;
                goProcess.Start();
                goProcess.WaitForExit();
                goProcess.Close();
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());                
            }

            StartTimer();
        }

    }
}
