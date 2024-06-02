# getLatestReleaseName
Get the name of the latest release of a GitHub repository.

## Who is this for?
For those who would like to check if the latest release is available

## install
```
go install github.com/UedaTakeyuki/getLatestReleaseName
```

## usage
### normal case
```
getLatestReleaseName yt-dlp yt-dlp
2024.05.27
```

### in case the repository doesn't have any release
```
getLatestReleaseName UedaTakeyuki mh-z19
2024/06/02 13:58:04 INFO No release

```

### in case the repository not found
```
getLatestReleaseName UedaTakeyuki mh-z1
2024/06/02 13:58:11 ERROR Not Found.

```

## How to use this as library in your application
```
import (
  "fmt"
	latest "github.com/UedaTakeyuki/getLatestReleaseName/latestReleaseName"
)

if name, err := latest.GetLatestReleaseName(user, repository); err != nil {
  slog.Error(err.Error())
} else {
	fmt.Println(name)
}

```

## History
2024.06.02 1.0.0 Created from https://github.com/UedaTakeyuki/myTrial/blob/main/goFileDownload/getLatestVersion.go
