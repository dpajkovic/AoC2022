#    Copyright (c) Miloš Rackov 2022
#    Distributed under the Boost Software License, Version 1.0.
#    (See accompanying file LICENSE or copy at
#    https://www.boost.org/LICENSE_1_0.txt)

$inputDay6 = Get-Content ./input.txt

for ($i = 3; $i -lt $inputDay6.Length; $i++)
{
    if(($inputDay6[($i - 3)..$i] | Select-Object -Unique).Count -eq 4)
    {
        break
    }
}

"Result for part 1: $($i+1)"

for ($i = 13; $i -lt $inputDay6.Length; $i++)
{
    if(($inputDay6[($i - 13)..$i] | Select-Object -Unique).Count -eq 14)
    {
        break
    }
}

"Result for part 2: $($i+1)"
