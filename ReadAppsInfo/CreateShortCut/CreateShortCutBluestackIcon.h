#pragma once

#include "resource.h"		

class CCreateShortCutBluestackIconApp : public CWinAppEx
{
public:
	CCreateShortCutBluestackIconApp();

	public:
	virtual BOOL InitInstance();

	DECLARE_MESSAGE_MAP()
};

extern CCreateShortCutBluestackIconApp theApp;