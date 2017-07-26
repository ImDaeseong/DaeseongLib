using Xamarin.Forms;

namespace App1
{
    public partial class App : Application
    {
        static bool bSplah = false;

        public App()
        {
            InitializeComponent();

            MainPage = new NavigationPage(new MainPage());

            /*
            if (bSplah)
            {
                MainPage = new NavigationPage(new MainPage());
            }
            else
            {
                bSplah = true;
                MainPage = new NavigationPage(new Splash());
            }
            */
        }

        protected override void OnStart()
        {
            // Handle when your app starts
        }

        protected override void OnSleep()
        {
            // Handle when your app sleeps
        }

        protected override void OnResume()
        {
            // Handle when your app resumes
        }
    }
}
