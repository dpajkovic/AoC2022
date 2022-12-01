
#    Copyright (c) Miloš Rackov 2022
#    Distributed under the Boost Software License, Version 1.0.
#    (See accompanying file LICENSE or copy at
#    https://www.boost.org/LICENSE_1_0.txt)

$inputDay1 = Get-Content ./input.txt -Raw

$elves = $inputDay1 -split "`n`n"

$calSums = @()

foreach($elf in $elves)
{
    $calories = $elf -split "`n"

    $sum = ($calories | Measure-Object -Sum).Sum
    $calSums += $sum

}

$maxSum = $calSums | Sort-Object -Descending -Top 1
$max3Sums = ($calSums | Sort-Object -Descending -Top 3 | Measure-Object -Sum).Sum

"Result for part 1: $maxSum"

"Result for part 2: $max3Sums"
