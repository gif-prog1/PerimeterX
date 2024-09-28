#include <windows.h>
#include <psapi.h>
#include <tchar.h>
#include <vector>
#include <string>
#include <iostream>

std::vector<DWORD> GetProcesses() {
    std::vector<DWORD> allProcesses;
    DWORD processes[1024], cbNeeded, cProcesses;

    if (EnumProcesses(processes, sizeof(processes), &cbNeeded)) {
        cProcesses = cbNeeded / sizeof(DWORD);

        for (unsigned int i = 0; i < cProcesses; i++) {
            if (processes[i] != 0) {
                HANDLE hProcess = OpenProcess(PROCESS_QUERY_INFORMATION | PROCESS_VM_READ, FALSE, processes[i]);
                if (hProcess != NULL) {
                    TCHAR szProcessName[MAX_PATH] = TEXT("<unknown>");
                    HMODULE hMod;
                    DWORD cbNeeded;

                    if (EnumProcessModules(hProcess, &hMod, sizeof(hMod), &cbNeeded)) {
                        GetModuleBaseName(hProcess, hMod, szProcessName, sizeof(szProcessName) / sizeof(TCHAR));
                    }

                    // Print process name and ID
                    std::wcout << "Process Name: " << szProcessName << ", Process ID: " << processes[i] << std::endl;

                    allProcesses.push_back(processes[i]);

                    CloseHandle(hProcess);
                }
            }
        }
    }

    return allProcesses;
}

int main() {
    GetProcesses();
    return 0;
}