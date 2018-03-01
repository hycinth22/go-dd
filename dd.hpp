/*
DD —— 驱动级键鼠操作模拟
http://www.ddxoft.com/help/
*/

#include <stdint.h>
#include <stdbool.h>
#include "windows.h"
HMODULE hMoudle = 0;

bool init() {
    const char* DLLPATH = "dd.dll";
    hMoudle = LoadLibrary(DLLPATH);
    return hMoudle!=NULL;
}
void* getProc(const char* procName) {
    if (hMoudle == 0) {
        return NULL;
    }
    return GetProcAddress(hMoudle, procName);
}

int32_t DD_btn(int32_t flag) {
    typedef int32_t(*ProcType)(int32_t);
    ProcType f = (ProcType)getProc("DD_btn");
    if (f == NULL) {
        return -1;
    }
    return f(flag);
}

int32_t DD_mov(int32_t x, int32_t y) {
    typedef int32_t(*ProcType)(int32_t, int32_t);
    ProcType f = (ProcType)getProc("DD_mov");
    if (f == NULL) {
        return -1;
    }
    return f(x, y);
}

int32_t DD_movR(int32_t dx, int32_t dy) {
    typedef int32_t(*ProcType)(int32_t, int32_t);
    ProcType f = (ProcType)getProc("DD_movR");
    if (f == NULL) {
        return -1;
    }
    return f(dx, dy);
}

int32_t DD_whl(int32_t flag) {
    typedef int32_t(*ProcType)(int32_t);
    ProcType f = (ProcType)getProc("DD_whl");
    if (f == NULL) {
        return -1;
    }
    return f(flag);
}

int32_t DD_key(int32_t key, int32_t status) {
    typedef int32_t(*ProcType)(int32_t, int32_t);
    ProcType f = (ProcType)getProc("DD_key");
    if (f == NULL) {
        return -1;
    }
    return f(key, status);
}

int32_t DD_str(const char* str){
    typedef int32_t(*ProcType)(const char* );
    ProcType f = (ProcType)getProc("DD_str");
    if (f == NULL) {
        return -1;
    }
    return f(str);
}
