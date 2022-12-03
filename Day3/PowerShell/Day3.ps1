#    Copyright (c) Miloš Rackov 2022
#    Distributed under the Boost Software License, Version 1.0.
#    (See accompanying file LICENSE or copy at
#    https://www.boost.org/LICENSE_1_0.txt)

$inputDay3 = Get-Content ./input.txt -Raw
# $inputDay3 = Get-Content ./testinput.txt -Raw

$rs = $inputDay3 -split "`n"
$letters1 = ""
foreach($r in $rs)
{
    $l = ""
    $i = 0
    while ($l -eq "")
    {
        if ($r.Substring($r.Length / 2).Contains($r[$i]))
        {
            $l = $r[$i]
        }
        $i++
    }
    $letters1 += $l
}
$sum1 = 0
for ($i = 0; $i -lt $letters1.Length; $i++)
{
    if ([int][char]$letters1[$i] -gt 95)
    {
        $sum1 += [int][char]$letters1[$i] - 96
    }
    else
    {
        $sum1 += [int][char]$letters1[$i] - 38
    }
}

$letters2 = ""
for ($i = 0; $i -lt $rs.Count; $i += 3)
{
    $r1 = $rs[$i]
    $r2 = $rs[$i + 1]
    $r3 = $rs[$i + 2]
    $l = ""
    $j = 0
    while ("" -eq $l)
    {
        if($r2.Contains($r1[$j]) -and $r3.Contains($r1[$j]))
        {
            $l = $r1[$j]
        }
        $j++
    }
    $letters2 += $l
}
$sum2 = 0
for ($i = 0; $i -lt $letters2.Length; $i++)
{
    if ([int][char]$letters2[$i] -gt 95)
    {
        $sum2 += [int][char]$letters2[$i] - 96
    }
    else
    {
        $sum2 += [int][char]$letters2[$i] - 38
    }
}
"Result for part 1: $sum1"

"Result for part 2: $sum2"
