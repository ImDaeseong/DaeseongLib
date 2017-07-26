using System;
using System.Threading.Tasks;
using Xamarin.Forms;
using Xamarin.Forms.Xaml;

namespace App1
{
    [XamlCompilation(XamlCompilationOptions.Compile)]
    public partial class MainPage : ContentPage
    {
        public MainPage()
        {
            InitializeComponent();

            NavigationPage.SetHasNavigationBar(this, false);
        }

        private async void TapGestureRecognizer_Tapped_localSearch(object sender, EventArgs e)
        {
            try
            {
                var content = new AllListView();
                ChangeContent.Content = content;
                SetTabTextColor(1);
                await AnimateStackLayout(localSearch);
            }
            catch
            {
            }
        }

        private async void TapGestureRecognizer_Tapped_PageSearch(object sender, EventArgs e)
        {
            try
            {
                var content = new PartListView();
                ChangeContent.Content = content;
                SetTabTextColor(2);
                await AnimateStackLayout(PageSearch);
            }
            catch
            {
            }
        }

        private async void TapGestureRecognizer_Tapped_FullSearch(object sender, EventArgs e)
        {
            try
            {
                var content = new SearchView();
                ChangeContent.Content = content;
                SetTabTextColor(3);
                await AnimateStackLayout(FullSearch);
            }
            catch
            {
            }
        }

        private async void TapGestureRecognizer_Tapped_SiteView(object sender, EventArgs e)
        {
            try
            {
                var content = new WebView();
                ChangeContent.Content = content;
                SetTabTextColor(4);
                await AnimateStackLayout(SiteView);
            }
            catch
            {
            }
        }

        private async Task<bool> AnimateStackLayout(StackLayout sl)
        {
            await sl.ScaleTo(0.9, 75, Easing.CubicOut);
            await sl.ScaleTo(1, 75, Easing.CubicIn);
            return true;
        }

        private void SetTabTextColor(int nIndex)
        {
            if (nIndex == 1)
            {
                tab1.TextColor = Color.Orange;
                tab2.TextColor = Color.FromHex("#5A297D");
                tab3.TextColor = Color.FromHex("#5A297D");
                tab4.TextColor = Color.FromHex("#5A297D");
            }
            else if (nIndex == 2)
            {
                tab1.TextColor = Color.FromHex("#5A297D");
                tab2.TextColor = Color.Orange;
                tab3.TextColor = Color.FromHex("#5A297D");
                tab4.TextColor = Color.FromHex("#5A297D");
            }
            else if (nIndex == 3)
            {
                tab1.TextColor = Color.FromHex("#5A297D");
                tab2.TextColor = Color.FromHex("#5A297D");
                tab3.TextColor = Color.Orange;
                tab4.TextColor = Color.FromHex("#5A297D");
            }
            else if (nIndex == 4)
            {
                tab1.TextColor = Color.FromHex("#5A297D");
                tab2.TextColor = Color.FromHex("#5A297D");
                tab3.TextColor = Color.FromHex("#5A297D");
                tab4.TextColor = Color.Orange;
            }
        }
    }
}