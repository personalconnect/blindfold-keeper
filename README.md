<p align="center">
  <img width="500" height="500" alt="Image" src="https://github.com/user-attachments/assets/11b193f0-b9ed-45e9-a7fa-923f0c8c47ae" />
</p>

This helper program supports the E2EE configuration for the [BlindFold Chrome Extension](https://chromewebstore.google.com/detail/blindfold/cmgjlocmnppfpknaipdfodjhbplnhimk?hl=ko&utm_source=ext_sidebar) **v1.0.2** by storing the device key in the Windows Credential Manager and the macOS Keychain.

## Download

Download the latest release from the [Releases page](https://github.com/personalconnect/dragpass-keeper/releases).

### Available Packages

- **macOS**: `dragpass-keeper-macos-arm64.pkg` (Apple Silicon)
- **Windows**: `dragpass-keeper.exe` (x64 installer)
- **Linux**:
  - `dragpass-keeper-x86_64.deb` (x86_64/amd64)
  - `dragpass-keeper-arm64.deb` (ARM64)

## Verifying Downloads

All release packages are signed with GPG for security. We strongly recommend verifying the integrity of downloaded files.

### 1. Import the Public Key

```bash
# Download and import the public key
curl https://raw.githubusercontent.com/personalconnect/dragpass-keeper/main/GPG_PUBLIC_KEY.asc | gpg --import
```

Or import manually from [GPG_KEYSPUBLIC_KEY.asc](GPG_PUBLIC_KEY.asc).

**Key Fingerprint**: `66DF 4017 8A5F 6F66 EAAF 318A 3FC4 1856 9192 8FDC`

### 2. Verify the Signature

```bash
# For macOS
gpg --verify dragpass-keeper-macos-arm64.pkg.sig dragpass-keeper-macos-arm64.pkg

# For Windows
gpg --verify dragpass-keeper.exe.sig dragpass-keeper.exe

# For Linux (x86_64)
gpg --verify dragpass-keeper-x86_64.deb.sig dragpass-keeper-x86_64.deb

# For Linux (ARM64)
gpg --verify dragpass-keeper-arm64.deb.sig dragpass-keeper-arm64.deb
```

You should see output like:
```
gpg: Good signature from "JinHyeok Hong <vjinhyeokv@gmail.com>" [ultimate]
```

## Output

### MacOS

- /Library/Application\ Support/Google/Chrome/NativeMessagingHosts/com.dragpass.keeper.json
- /Library/Application Support/DragPass/mac-dragpass-keeper

### Windows (64bit)

- C:\Program Files\DragPass
  - com.dragpass.keeper.json
  - unins000.exe
  - unins000.dat
  - dragpass-keeper.exe

### Windows (32bit)

- C:\Program Files (x86)\DragPass
  - com.dragpass.keeper.json
  - unins000.exe
  - unins000.dat
  - dragpass-keeper.exe