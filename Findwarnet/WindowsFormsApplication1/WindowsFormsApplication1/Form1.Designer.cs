namespace WindowsFormsApplication1
{
    partial class frmMain
    {
        /// <summary>
        /// 필수 디자이너 변수입니다.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// 사용 중인 모든 리소스를 정리합니다.
        /// </summary>
        /// <param name="disposing">관리되는 리소스를 삭제해야 하면 true이고, 그렇지 않으면 false입니다.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form 디자이너에서 생성한 코드

        /// <summary>
        /// 디자이너 지원에 필요한 메서드입니다. 
        /// 이 메서드의 내용을 코드 편집기로 수정하지 마세요.
        /// </summary>
        private void InitializeComponent()
        {
            this.lstView = new System.Windows.Forms.ListView();
            this.txtPhone = new System.Windows.Forms.TextBox();
            this.txtMapurl = new System.Windows.Forms.TextBox();
            this.txtAuthorUrl = new System.Windows.Forms.TextBox();
            this.SuspendLayout();
            // 
            // lstView
            // 
            this.lstView.Location = new System.Drawing.Point(12, 13);
            this.lstView.Name = "lstView";
            this.lstView.Size = new System.Drawing.Size(1338, 659);
            this.lstView.TabIndex = 3;
            this.lstView.UseCompatibleStateImageBehavior = false;
            this.lstView.ColumnClick += new System.Windows.Forms.ColumnClickEventHandler(this.lstView_ColumnClick);
            this.lstView.DrawColumnHeader += new System.Windows.Forms.DrawListViewColumnHeaderEventHandler(this.lstView_DrawColumnHeader);
            this.lstView.DrawItem += new System.Windows.Forms.DrawListViewItemEventHandler(this.lstView_DrawItem);
            this.lstView.DrawSubItem += new System.Windows.Forms.DrawListViewSubItemEventHandler(this.lstView_DrawSubItem);
            this.lstView.SelectedIndexChanged += new System.EventHandler(this.lstView_SelectedIndexChanged);
            // 
            // txtPhone
            // 
            this.txtPhone.Location = new System.Drawing.Point(12, 697);
            this.txtPhone.Name = "txtPhone";
            this.txtPhone.Size = new System.Drawing.Size(239, 21);
            this.txtPhone.TabIndex = 4;
            // 
            // txtMapurl
            // 
            this.txtMapurl.Location = new System.Drawing.Point(267, 697);
            this.txtMapurl.Name = "txtMapurl";
            this.txtMapurl.Size = new System.Drawing.Size(525, 21);
            this.txtMapurl.TabIndex = 5;
            // 
            // txtAuthorUrl
            // 
            this.txtAuthorUrl.Location = new System.Drawing.Point(804, 697);
            this.txtAuthorUrl.Name = "txtAuthorUrl";
            this.txtAuthorUrl.Size = new System.Drawing.Size(549, 21);
            this.txtAuthorUrl.TabIndex = 6;
            // 
            // frmMain
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(7F, 12F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(1364, 738);
            this.Controls.Add(this.txtAuthorUrl);
            this.Controls.Add(this.txtMapurl);
            this.Controls.Add(this.txtPhone);
            this.Controls.Add(this.lstView);
            this.MaximizeBox = false;
            this.MinimizeBox = false;
            this.Name = "frmMain";
            this.Text = "Form1";
            this.Load += new System.EventHandler(this.Form1_Load);
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion
        private System.Windows.Forms.ListView lstView;
        private System.Windows.Forms.TextBox txtPhone;
        private System.Windows.Forms.TextBox txtMapurl;
        private System.Windows.Forms.TextBox txtAuthorUrl;
    }
}

