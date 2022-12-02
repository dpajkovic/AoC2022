
#    Copyright (c) Miloš Rackov 2022
#    Distributed under the Boost Software License, Version 1.0.
#    (See accompanying file LICENSE or copy at
#    https://www.boost.org/LICENSE_1_0.txt)

$inputDay2 = Get-Content ./input.txt -Raw
# $inputDay2 = Get-Content ./testinput.txt -Raw

$lup1 = @{}
$lup1["A X"] = 4; $lup1["A Y"] = 8; $lup1["A Z"] = 3
$lup1["B X"] = 1; $lup1["B Y"] = 5; $lup1["B Z"] = 9
$lup1["C X"] = 7; $lup1["C Y"] = 2; $lup1["C Z"] = 6

$lup2 = @{}
$lup2["A X"] = 0 + 3; $lup2["A Y"] = 3 + 1; $lup2["A Z"] = 6 + 2
$lup2["B X"] = 0 + 1; $lup2["B Y"] = 3 + 2; $lup2["B Z"] = 6 + 3
$lup2["C X"] = 0 + 2; $lup2["C Y"] = 3 + 3; $lup2["C Z"] = 6 + 1

$rounds = $inputDay2 -split "`n"

$sum1 = 0
$sum2 = 0

foreach($round in $rounds)
{
    $sum1 += $lup1[$round]
    $sum2 += $lup2[$round]
}

"Result for part 1: $sum1"

"Result for part 2: $sum2"
