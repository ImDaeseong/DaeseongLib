using Android.App;
using Android.Content.PM;
using Android.OS;
using Android.Support.V7.App;

namespace App1.Droid
{
    /*
    [Activity(Label = "App1", Icon = "@drawable/icon", Theme = "@style/Theme.Splash", MainLauncher = true, NoHistory = true, ConfigurationChanges = ConfigChanges.ScreenSize | ConfigChanges.Orientation, LaunchMode = LaunchMode.SingleTop)]
    public class SplashActivity : AppCompatActivity
    {
        protected override void OnCreate(Bundle bundle)
        {
            base.OnCreate(bundle);

            StartActivity(typeof(MainActivity));

            Finish();
        }

        public override void OnBackPressed()
        {
            return;
        }
    };
    */

    //[Activity(Label = "App1", Icon = "@drawable/icon", Theme = "@style/MainTheme", MainLauncher = false, ConfigurationChanges = ConfigChanges.ScreenSize | ConfigChanges.Orientation)]
    [Activity(Label = "App1", Icon = "@drawable/icon", Theme = "@style/Theme.Splash", MainLauncher = true, ConfigurationChanges = ConfigChanges.ScreenSize | ConfigChanges.Orientation)]
    public class MainActivity : global::Xamarin.Forms.Platform.Android.FormsAppCompatActivity
    {
        protected override void OnCreate(Bundle bundle)
        {
            base.OnCreate(bundle);

            global::Xamarin.Forms.Forms.Init(this, bundle);
            global::Xamarin.FormsMaps.Init(this, bundle);

            LoadApplication(new App());
        }
    }
}

