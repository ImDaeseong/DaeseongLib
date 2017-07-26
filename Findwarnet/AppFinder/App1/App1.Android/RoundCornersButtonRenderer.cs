using Android.Graphics.Drawables;
using App1;
using App1.Droid;
using Xamarin.Forms;
using Xamarin.Forms.Platform.Android;

[assembly: ExportRenderer(typeof(RoundCornersButton), typeof(RoundCornersButtonRenderer))]
namespace App1.Droid
{
    public class RoundCornersButtonRenderer : ButtonRenderer
    {
        protected override void OnElementChanged(ElementChangedEventArgs<Xamarin.Forms.Button> e)
        {
            base.OnElementChanged(e);
            if (e.NewElement != null)
            {
                //Subscribe to the events and stuff
            }
            else if (e.OldElement != null)
            {
                //Unsubscribe from events
            }

            if (Control != null)
            {
                GradientDrawable gradientDrawable = new GradientDrawable();
                gradientDrawable.SetShape(ShapeType.Rectangle);
                gradientDrawable.SetColor(Element.BackgroundColor.ToAndroid());
                gradientDrawable.SetStroke(4, Element.BorderColor.ToAndroid());
                gradientDrawable.SetCornerRadius(38.0f);
                Control.SetBackground(gradientDrawable);
            }
        }
    }
}