using System;
using System.Windows.Forms;
using System.Data.SQLite;
using System.Data;
using System.Drawing;

namespace WindowsFormsApplication1
{
    public partial class frmMain : Form
    {       
        clsDBConnect DBConn = clsDBConnect.getInstance;
        clsComm CFunc = clsComm.getInstance;
        
        public frmMain()
        {
            InitializeComponent();
        }
        
        private void Form1_Load(object sender, EventArgs e)
        {
            if (DBConn.DbConnection())
            {
                Console.WriteLine("connected");
            }
            else
            {
                Console.WriteLine("not connected");                
            }

            //테이블 생성
            //DBConn.createTableGooglePlaces();

            LoadData();
        }

        private void LoadData()
        {
            lstView.View = View.Details;
            lstView.GridLines = true;
            lstView.FullRowSelect = true;
            lstView.HeaderStyle = ColumnHeaderStyle.Clickable;
            lstView.CheckBoxes = true;
            lstView.OwnerDraw = true;

            lstView.Columns.Add("", 25, HorizontalAlignment.Left);
            lstView.Columns.Add("번호", 0, HorizontalAlignment.Left);
            lstView.Columns.Add("Pc방", 180, HorizontalAlignment.Left);
            lstView.Columns.Add("위도", 0, HorizontalAlignment.Left);
            lstView.Columns.Add("경도", 0, HorizontalAlignment.Left);
            lstView.Columns.Add("place_id", 0, HorizontalAlignment.Left);
            lstView.Columns.Add("주소", 200, HorizontalAlignment.Left);
            lstView.Columns.Add("전화번호", 100, HorizontalAlignment.Left);
            lstView.Columns.Add("인터넷전화번호", 110, HorizontalAlignment.Left);
            lstView.Columns.Add("Url", 150, HorizontalAlignment.Left);
            lstView.Columns.Add("Authorname", 100, HorizontalAlignment.Left);
            lstView.Columns.Add("Authorurl", 110, HorizontalAlignment.Left);
            lstView.Columns.Add("Profilephotourl", 100, HorizontalAlignment.Left);
            lstView.Columns.Add("Relativetimedescription", 100, HorizontalAlignment.Left);
            lstView.Columns.Add("TextDesc", 150, HorizontalAlignment.Left);

            string sQuery = string.Format("SELECT * FROM tbPlaceId");
            SQLiteDataReader r = DBConn.GetSelectQuery(sQuery);
            while (r.Read())
            {
                String s1 = r["rIndex"].ToString();
                String s2 = r["Name"].ToString();
                String s3 = r["Lat"].ToString();
                String s4 = r["Lng"].ToString();
                String s5 = r["place_id"].ToString();
                String s6 = r["Address"].ToString();
                String s7 = r["phonenumber"].ToString();
                String s8 = r["Internatphonenumber"].ToString();
                String s9 = r["Url"].ToString();
                String s10 = r["Authorname"].ToString();
                String s11 = r["Authorurl"].ToString();
                String s12 = r["Profilephotourl"].ToString();
                String s13 = r["Relativetimedescription"].ToString();
                String s14 = r["TextDesc"].ToString();

                ListViewItem item = new ListViewItem();
                item.Text = "";
                item.SubItems.Add(s1);
                item.SubItems.Add(s2);
                item.SubItems.Add(s3);
                item.SubItems.Add(s4);
                item.SubItems.Add(s5);
                item.SubItems.Add(s6);
                item.SubItems.Add(s7);
                item.SubItems.Add(s8);
                item.SubItems.Add(s9);
                item.SubItems.Add(s10);
                item.SubItems.Add(s11);
                item.SubItems.Add(s12);
                item.SubItems.Add(s13);
                item.SubItems.Add(s14);
                lstView.Items.Add(item);
            }

            this.Text = string.Format("현재 총{0}개의 PC방 정보가 등록되어 있습니다.", lstView.Items.Count);
        }

        private void InitChkBox()
        {
            for (int i = 0; i < lstView.Items.Count; i++)
                lstView.Items[i].Checked = false;
        }

        private void lstView_DrawColumnHeader(object sender, DrawListViewColumnHeaderEventArgs e)
        {
            if (e.ColumnIndex == 0)
            {
                e.DrawBackground();
                bool value = false;
                try
                {
                    value = Convert.ToBoolean(e.Header.Tag);
                }
                catch (Exception)
                {
                }
                CheckBoxRenderer.DrawCheckBox(e.Graphics,
                    new Point(e.Bounds.Left + 4, e.Bounds.Top + 4),
                    value ? System.Windows.Forms.VisualStyles.CheckBoxState.CheckedNormal :
                    System.Windows.Forms.VisualStyles.CheckBoxState.UncheckedNormal);
            }
            else
            {
                e.DrawDefault = true;
            }
        }

        private void lstView_DrawItem(object sender, DrawListViewItemEventArgs e)
        {
            e.DrawDefault = true;
        }

        private void lstView_DrawSubItem(object sender, DrawListViewSubItemEventArgs e)
        {
            e.DrawDefault = true;
        }
        
        private void lstView_ColumnClick(object sender, ColumnClickEventArgs e)
        {
            if (e.Column == 0)
            {
                bool value = false;
                try
                {
                    value = Convert.ToBoolean(lstView.Columns[e.Column].Tag);
                }
                catch (Exception)
                {
                }
                lstView.Columns[e.Column].Tag = !value;
                foreach (ListViewItem item in lstView.Items)
                    item.Checked = !value;

                lstView.Invalidate();
            }
        }       

        private void lstView_SelectedIndexChanged(object sender, EventArgs e)
        {
            if (lstView.SelectedItems.Count > 0)
            {
                string rIndex = lstView.SelectedItems[0].SubItems[1].Text.ToString();
                string Name = lstView.SelectedItems[0].SubItems[2].Text.ToString();
                string Lat = lstView.SelectedItems[0].SubItems[3].Text.ToString();
                string Lng = lstView.SelectedItems[0].SubItems[4].Text.ToString();
                string place_id = lstView.SelectedItems[0].SubItems[5].Text.ToString();
                string Address = lstView.SelectedItems[0].SubItems[6].Text.ToString();
                string phonenumber = lstView.SelectedItems[0].SubItems[7].Text.ToString();
                string Internatphonenumber = lstView.SelectedItems[0].SubItems[8].Text.ToString();
                string Url = lstView.SelectedItems[0].SubItems[9].Text.ToString();
                string Authorname = lstView.SelectedItems[0].SubItems[10].Text.ToString();
                string Authorurl = lstView.SelectedItems[0].SubItems[11].Text.ToString();
                string Profilephotourl = lstView.SelectedItems[0].SubItems[12].Text.ToString();
                string Relativetimedescription = lstView.SelectedItems[0].SubItems[13].Text.ToString();
                string TextDesc = lstView.SelectedItems[0].SubItems[14].Text.ToString();

                txtPhone.Text = phonenumber;
                txtMapurl.Text = Url;
                txtAuthorUrl.Text = Authorurl;
            }
            else
            {
                txtPhone.Text = "";
                txtMapurl.Text = "";
                txtAuthorUrl.Text = "";
            }
        }
        
    }
}
