using System.Threading.Tasks;
using Xamarin.Forms;
using Xamarin.Forms.Xaml;

namespace App1
{
    [XamlCompilation(XamlCompilationOptions.Compile)]
    public partial class Splash : ContentPage
    {
        public Splash()
        {
            InitializeComponent();

            NavigationPage.SetHasNavigationBar(this, false);
            Device.BeginInvokeOnMainThread(() => DelayedShowFirstView());
        }

        private async Task DelayedShowFirstView()
        {
            await Task.Delay(1);
            await Navigation.PushAsync(new MainPage());
            Navigation.RemovePage(this);
        }

        protected override bool OnBackButtonPressed()
        {
            return true;
        }
    }
}