
#    Copyright (c) Miloš Rackov 2022
#    Distributed under the Boost Software License, Version 1.0.
#    (See accompanying file LICENSE or copy at
#    https://www.boost.org/LICENSE_1_0.txt)

function Get-EnvVals
{
    $pairs = Get-Content ./AoC.env
    foreach($pair in $pairs)
    {
        $name, $value = $pair -split "="
        Set-Content Env:\$name $value
    }
}

Get-EnvVals

$session = [Microsoft.PowerShell.Commands.WebRequestSession]::new()
$cookie = [System.Net.Cookie]::new($env:SESSION_COOKIE_NAME, $env:SESSION_COOKIE_VALUE, $env:SESSION_COOKIE_PATH, $env:SESSION_COOKIE_DOMAIN)
$session.Cookies.Add($cookie)

$leaderBoard = Invoke-RestMethod -Method Get -Uri $env:SESSION_URL -WebSession $session -ContentType "application/json"

$topList = $leaderBoard.members.psobject.Properties.value | Sort-Object -Property local_score -Top 10 -Descending | Select-Object -Property id, name, stars, local_score

$sortedList = @()

for ($i = 0; $i -lt $topList.Count; $i++)
{
    $sortedList += [PSCustomObject]@{
        Order         = $i + 1
        Name          = $topList[$i].name
        Stars         = $topList[$i].stars
        "Local Score" = $topList[$i].local_score
    }
}

$sortedList | Format-Table

Import-Module FormatMarkdownTable

$md = $sortedList | Format-MarkdownTableTableStyle -ShowMarkdown -HideStandardOutput -Property Order, Name, Stars, "Local Score" -DoNotCopyToClipboard

$md = $md.Replace("|:--|:--|:--|:--|", "|--:|:--|--:|--:|")
$day = (Get-Date).Day
$myStars = $leaderBoard.members.$env:SESSION_MEMBER_ID.stars

$readme = @'
# Advent of Code 2022

![day](https://img.shields.io/badge/day%20📅-{0}-blue)
![stars](https://img.shields.io/badge/stars%20⭐-{1}x-yellow)

[Advent of Code](https://adventofcode.com)

## Local Leaderboard JPMC

{2}
'@ -f $day, $myStars, $md

$readme | Out-File ./README.md
