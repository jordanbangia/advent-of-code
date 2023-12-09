
import math


times = [47847467]
distance = [207139412091014]


w_prod = 1

for time, distance in zip(times, distance):

    a = -1
    b = time
    c = -distance

    root_1 = (-1 * b + math.sqrt(b*b - 4 * a * c)) / 2*a
    root_2 = (-1 * b - math.sqrt(b*b - 4 * a * c)) / 2*a

    print(root_1, root_2)

    root_1 = math.ceil(root_1) if math.ceil(root_1) > root_1 else root_1 + 1
    root_2 = math.floor(root_2) if math.floor(root_2) < root_2 else root_2 - 1 

    ways_to_win = root_2 - root_1 + 1
    print(root_1, root_2, ways_to_win)
    w_prod *= ways_to_win

print(w_prod)
    



'''

y = distance travelled
y = (t - x)*x

f(0) = 0*(0 - t) = 0
f(1) = 1* (7 - 1) = 6
f(2) = 2 * (7-2) = 2* 5 = 10
f(7) = 7 * (7 - 7) = 0


Ask is to find x s.t (tx - x*x) > d

tx - x^2 > d
tx - x^2 -d > 0


tx - x^2 - d = 0
x^2 - tx + d = 0

t +/- sqrt(t^2 - 4d)

3.5 +/- 1/2 * (49 - 4*9)^0.5
= 3.5 +/- 0.5 * (49 - 36)^0.5
= 3.5 +/- 0.5 * (13)^0.5


'''