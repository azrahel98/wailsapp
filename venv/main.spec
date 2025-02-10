# -*- mode: python ; coding: utf-8 -*-

from PyInstaller.utils.hooks import collect_all

datas = [('.env', '.')]
binaries = []
hiddenimports = [
    'mysql.connector',
    'langchain',
    'langchain_core',
    'langchain_community',
    'langchain_openai',
    'dotenv',
    'pydantic',
    'pydantic.deprecated',
    'pydantic.deprecated.decorator',
    'pydantic._internal',
    'pydantic_core',
    'pydantic.version',
    'pydantic.warnings',
    'typing_extensions'
]

# Collect all packages
for package in ['pydantic', 'langchain_openai', 'langchain_core']:
    collect_result = collect_all(package)
    datas.extend(collect_result[0])
    binaries.extend(collect_result[1])
    hiddenimports.extend(collect_result[2])

a = Analysis(
    ['maincd.py'],
    pathex=[],
    binaries=binaries,
    datas=datas,
    hiddenimports=hiddenimports,
    hookspath=[],
    hooksconfig={},
    runtime_hooks=[],
    excludes=[],
    win_no_prefer_redirects=False,
    win_private_assemblies=False,
    noarchive=False
)

pyz = PYZ(a.pure, a.zipped_data)

exe = EXE(
    pyz,
    a.scripts,
    a.binaries,
    a.zipfiles,
    a.datas, 
    [],
    name='consulta',
    debug=False,
    bootloader_ignore_signals=False,
    strip=False,
    upx=True,
    upx_exclude=[],
    runtime_tmpdir=None,
    console=True,
    disable_windowed_traceback=False,
    argv_emulation=False,
    target_arch=None,
    codesign_identity=None,
    entitlements_file=None,
)