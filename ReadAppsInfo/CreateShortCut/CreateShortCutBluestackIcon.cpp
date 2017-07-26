#include "stdafx.h"
#include "CreateShortCutBluestackIcon.h"

BEGIN_MESSAGE_MAP(CCreateShortCutBluestackIconApp, CWinAppEx)
END_MESSAGE_MAP()

CCreateShortCutBluestackIconApp::CCreateShortCutBluestackIconApp()
{	
}

CCreateShortCutBluestackIconApp theApp;

BOOL CCreateShortCutBluestackIconApp::InitInstance()
{
	CoInitialize(NULL);

	CCommandLineInfo cmdInfo;
	ParseCommandLine(cmdInfo);
	CString strParams = cmdInfo.m_strFileName;
	if(strParams == "")
	{
		Bluestacks m_objApps;
		BOOL bLoad = m_objApps.ReadJson();
		if(bLoad)
			m_objApps.CreateAppIconList();
	}
	else
	{
		CString strValue1, strValue2, strValue3, strValue4, strValue5;
		AfxExtractSubString(strValue1, strParams, 0, '|');
		AfxExtractSubString(strValue2, strParams, 1, '|'); 
		AfxExtractSubString(strValue3, strParams, 2, '|'); 
		AfxExtractSubString(strValue4, strParams, 3, '|');
		AfxExtractSubString(strValue5, strParams, 4, '|');
		Bluestacks m_objApps;
		m_objApps.CreateAppIconList(strValue1, strValue2, strValue3, strValue4, strValue5);
	}
	return FALSE;
}
