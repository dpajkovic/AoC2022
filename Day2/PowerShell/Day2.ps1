
#    Copyright (c) Miloš Rackov 2022
#    Distributed under the Boost Software License, Version 1.0.
#    (See accompanying file LICENSE or copy at
#    https://www.boost.org/LICENSE_1_0.txt)

$inputDay2 = Get-Content ./input.txt -Raw
# $inputDay2 = Get-Content ./testinput.txt -Raw

$lut1 = @{}
$lut1["A X"] = 3 + 1; $lut1["A Y"] = 6 + 2; $lut1["A Z"] = 0 + 3
$lut1["B X"] = 0 + 1; $lut1["B Y"] = 3 + 2; $lut1["B Z"] = 6 + 3
$lut1["C X"] = 6 + 1; $lut1["C Y"] = 0 + 2; $lut1["C Z"] = 3 + 3

$lut2 = @{}
$lut2["A X"] = 0 + 3; $lut2["A Y"] = 3 + 1; $lut2["A Z"] = 6 + 2
$lut2["B X"] = 0 + 1; $lut2["B Y"] = 3 + 2; $lut2["B Z"] = 6 + 3
$lut2["C X"] = 0 + 2; $lut2["C Y"] = 3 + 3; $lut2["C Z"] = 6 + 1

$rounds = $inputDay2 -split "`n"

$sum1 = 0
$sum2 = 0

foreach($round in $rounds)
{
    $sum1 += $lut1[$round]
    $sum2 += $lut2[$round]
}

"Result for part 1: $sum1"

"Result for part 2: $sum2"
