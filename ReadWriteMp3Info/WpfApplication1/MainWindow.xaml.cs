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

namespace WpfApplication1
{
    /// <summary>
    /// MainWindow.xaml에 대한 상호 작용 논리
    /// </summary>
    public partial class MainWindow : Window
    {
        private string strFileName = "";

        public MainWindow()
        {
            InitializeComponent();

            Init();
        }

        private void Init(bool bIsEdit = false)
        {
            lblDrag.Margin = new Thickness(23, 62, 25, 55);

            if (bIsEdit)
            {
                lblDrag.Visibility = Visibility.Hidden;
                pnlText.Visibility = Visibility.Visible;
                
                
            }
            else
            {
                lblDrag.Visibility = Visibility.Visible;
                pnlText.Visibility = Visibility.Hidden;
            }
        }

        private void Window_Loaded(object sender, RoutedEventArgs e)
        {
            this.AllowDrop = true;
        }

        private void Window_Closed(object sender, EventArgs e)
        {

        }

        private void Window_MouseLeftButtonDown(object sender, MouseButtonEventArgs e)
        {
            base.OnMouseLeftButtonDown(e);

            this.DragMove();
        }

        private void Window_DragOver(object sender, DragEventArgs e)
        {
            if (!e.Data.GetDataPresent(DataFormats.FileDrop))
            {
                e.Effects = DragDropEffects.Move;
                e.Handled = true;
            }
        }

        string GetFileExtName(string strFilename)
        {
            int nPos = strFilename.LastIndexOf('.');
            int nLength = strFilename.Length;
            if (nPos < nLength)
                return strFilename.Substring(nPos + 1, (nLength - nPos) - 1);
            return string.Empty;
        }

        private void Window_Drop(object sender, DragEventArgs e)
        {
            strFileName = "";
            if (e.Data.GetDataPresent(DataFormats.FileDrop))
            {               
                string[] files = (string[])e.Data.GetData(DataFormats.FileDrop);
                foreach (string file in files)
                {
                    if (GetFileExtName(file).ToLower() == "mp3")
                    {
                        strFileName = file;
                        break;
                    }
                }
            }

            if (strFileName != "")
            {
                Init(true);
            }                      
        }

        private void SendFlag()
        {
            string strTitle = Title.Text;
            string strArtist = Artist.Text;
            string strAlbum = Album.Text;
            string strYear = Year.Text;
            string strComment = Comment.Text;

            //golang flag 들을경우 공백이 잘린다 그래서 공백을 <br> 변경후
            strFileName = strFileName.Replace(" ", "<br>");
            strTitle = strTitle.Replace(" ", "<br>");
            strArtist = strArtist.Replace(" ", "<br>");
            strAlbum = strAlbum.Replace(" ", "<br>");
            strYear = strYear.Replace(" ", "<br>");
            strComment = strComment.Replace(" ", "<br>");

            System.Diagnostics.Process goProcess = new System.Diagnostics.Process();
            goProcess.StartInfo.FileName = "D:\\Daeseonglib.exe";
            goProcess.StartInfo.Arguments = " -Path=" + strFileName + " -title=" + strTitle + " -artist=" + strArtist + " -album=" + strAlbum + " -year=" + strYear + " -comment=" + strComment;
            goProcess.StartInfo.WindowStyle = System.Diagnostics.ProcessWindowStyle.Hidden;
            goProcess.Start();
            goProcess.WaitForExit();
            goProcess.Close();

            Init(false);

            Title.Text = "";
            Artist.Text = "";
            Album.Text = "";
            Year.Text = "";
            Comment.Text = "";
        }

        private void SendArgs()
        {
            string strTitle = Title.Text;
            string strArtist = Artist.Text;
            string strAlbum = Album.Text;
            string strYear = Year.Text;
            string strComment = Comment.Text;

            //golang flag 들을경우 공백이 잘린다 그래서 공백을 <br> 변경후
            strFileName = strFileName.Replace(" ", "<br>");
            strTitle = strTitle.Replace(" ", "<br>");
            strArtist = strArtist.Replace(" ", "<br>");
            strAlbum = strAlbum.Replace(" ", "<br>");
            strYear = strYear.Replace(" ", "<br>");
            strComment = strComment.Replace(" ", "<br>");

            System.Diagnostics.Process goProcess = new System.Diagnostics.Process();
            goProcess.StartInfo.FileName = "D:\\Daeseonglib.exe";
            goProcess.StartInfo.Arguments = " " + strFileName + " " + strTitle + " " + strArtist + " " + strAlbum + " " + strYear + " " + strComment;
            goProcess.StartInfo.WindowStyle = System.Diagnostics.ProcessWindowStyle.Hidden;
            goProcess.Start();
            goProcess.WaitForExit();
            goProcess.Close();

            Init(false);

            Title.Text = "";
            Artist.Text = "";
            Album.Text = "";
            Year.Text = "";
            Comment.Text = "";
        }

        private void Button_Click(object sender, RoutedEventArgs e)
        {
            //SendArgs();

            SendFlag();           
        }

    }
}
