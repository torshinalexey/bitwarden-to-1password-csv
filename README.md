# Bitwarden to 1Password csv converter

This cli tool helps to convert exported logins from [Bitwarden](https://bitwarden.com) to [1Password](https://1password.com) compatible csv format.

1. [Export Bitwarden vault to csv file](https://bitwarden.com/help/article/export-your-data/#export-a-personal-vault)
2. Use exported csv file as input
3. [Import result to 1Password](https://support.1password.com/import-1password-com/)

## Options

**`-h`** usage info

**`-i`** path to csv file with logins exported from  (default "in.csv")

**`-o`** csv file name which will has logins in format accepted by  (default "out.csv")

## Example

```shell
bwto1pcsv_macos -i bw_logins.csv -o 1p_logins.csv
```