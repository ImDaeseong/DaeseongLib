using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Xamarin.Forms;
using System.Diagnostics;
using System.Threading;

namespace App1
{
    public partial class MainPage : ContentPage
    {
        private IDownloadPath dPath;

        public MainPage()
        {
            InitializeComponent();

            dPath = DependencyService.Get<IDownloadPath>();
        }

        private async void btnDownloadText_Clicked(object sender, EventArgs e)
        {
            string sUrl = "http://IP:8080/Daeseong";

            var dl = new HttpClientService();
            var sContext = await dl.DownloadTextAsync(sUrl);
           
            lblDownloadInfo.Text = sContext;
        }

        private async void btnDownloadFile_Clicked(object sender, EventArgs e)
        {
            string sUrl = "http://IP:8080/Daeseong/b.mp3";

            var dl = new HttpClientService();
            var bDown = await dl.DownloadFileAsync(sUrl);
            
            lblDownloadInfo.Text = bDown.ToString();            
        }

        private void btnDownloadProgressFile_Clicked(object sender, EventArgs e)
        {            
            var dl = new HttpClientProgressService(pbProgressBar);
            
            Task.Run(async () =>
            {
                var cancellationToken = new CancellationTokenSource();
                await dl.DownloadFileAsync(
                     "http://IP:8080/Daeseong/b.mp3",
                    cancellationToken.Token,
                    (message) =>
                    {
                        Device.BeginInvokeOnMainThread(async () =>
                        {
                            await this.DisplayAlert(
                            "Download Failed",
                            message,
                            "Ok");
                        });
                    },
                    () =>
                    {
                        Device.BeginInvokeOnMainThread(async () =>
                        {
                            await this.DisplayAlert(
                            "Download Complete",
                            "Download has completed",
                            "Ok");
                        });
                    });
            });
                        
        }
    }
}
