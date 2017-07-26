using Xamarin.Forms;
using Xamarin.Forms.Xaml;

namespace App1
{
    [XamlCompilation(XamlCompilationOptions.Compile)]
    public partial class WebView : ContentView
    {
        public WebView()
        {
            InitializeComponent();

            WebVeiw.Source = "http://m.naver.com";
        }

        private void WebVeiw_Navigated(object sender, WebNavigatedEventArgs e)
        {
            frame1.IsVisible = false;
            frame2.IsVisible = true;
        }

        private void WebVeiw_Navigating(object sender, WebNavigatingEventArgs e)
        {
            frame1.IsVisible = true;
            frame2.IsVisible = false;
        }
    }
}