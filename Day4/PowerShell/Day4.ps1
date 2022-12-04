#    Copyright (c) Miloš Rackov 2022
#    Distributed under the Boost Software License, Version 1.0.
#    (See accompanying file LICENSE or copy at
#    https://www.boost.org/LICENSE_1_0.txt)

$inputDay4 = Get-Content ./input.txt -Raw
# $inputDay4 = Get-Content ./testinput.txt -Raw

$counter1 = 0
$counter2 = 0
$pairs = $inputDay4 -split "`n"

foreach($pair in $pairs)
{
    $pair -match '(\d+)-(\d+),(\d+)-(\d+)' | Out-Null
    if((([int]$Matches[1] - [int]$Matches[3]) * ([int]$Matches[2] - [int]$Matches[4])) -le 0)
    {
        $counter1++
    }
    if(([int]$Matches[2] -ge [int]$Matches[3]) -and ([int]$Matches[4] -ge [int]$Matches[1]) )
    {
        $counter2++
    }

}
"Result for part 1: $counter1"

"Result for part 2: $counter2"
