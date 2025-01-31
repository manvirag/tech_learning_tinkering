Note: Slicing in array and slice is constant time , since it only change header.  Copy create copy always. Rule of Thumb => append() reallocates a new underlying array only if the slice’s capacity is exceeded. Otherwise, it modifies the original array in-place.

<img width="606" alt="image" src="https://github.com/user-attachments/assets/9296a915-d515-4c8a-b323-723c8df523f2" />


1. Channel Cheat sheet for error outing: https://blog.devtrovert.com/p/go-channels-explained-more-than-just

2. Hands On cheat sheet for golang: https://devhints.io/go

3. ![image](https://github.com/manvirag982/tech_learning_tinkering/assets/54881553/5cb9775d-3604-46b5-8666-fbb579b5689e)


4. <img width="891" alt="image" src="https://github.com/user-attachments/assets/69d8919c-f83d-4cb7-bd8d-a76c4ca74893" />
<br/>
<br/>

5. <img width="590" alt="image" src="https://github.com/user-attachments/assets/c9d62f7d-5359-4438-8de9-09cff8494388" />

  - Pass By Value => it directly copy the data, now modifying( like arr[0] = 2 ) and reassigning/copying ( arr = append(arr,3) ) won't affect orignal 
  - Pass By Value ( Pointer ) => all these thing work on pointer , pointer point the data of these , when passed in function then it copy the pointer and data is same shared. ![image](https://github.com/user-attachments/assets/1e867c9e-5442-426e-a396-c146b19258bc)
  - So it mean in above case, => when we modifying then it will change, but when we copying or reassgning then that pointer will shift to other data and now no longer pointing to shared data and won't affect original. append operation create copy.
    
<br/>
<br/>

6. <img width="599" alt="image" src="https://github.com/user-attachments/assets/c9d0bd0d-709c-4d26-9ba8-d3fb69263d38" />

<br/>
<br/>

7. <img width="592" alt="image" src="https://github.com/user-attachments/assets/27d1c0d8-2405-4139-bd51-54dbdc3490e2" />

<br/>
<br/>

8. Cheat Sheet
 
![image](https://github.com/user-attachments/assets/8fb5d984-02bd-4d86-b643-24370ee96d08)

<br/>
<br/>

9. STL

<img width="601" alt="image" src="https://github.com/user-attachments/assets/a636e0c2-bf0e-4836-bf6b-dd41fd3d2124" />

<img width="601" alt="image" src="https://github.com/user-attachments/assets/9505241d-e9cb-44dd-9652-0c63f71d5aaa" />

<img width="601" alt="image" src="https://github.com/user-attachments/assets/c27ee1d1-aa99-4341-9126-a042532ec004" />

<img width="601" alt="image" src="https://github.com/user-attachments/assets/82f3858b-24df-4fd6-bbd2-0209e2bfcbb2" />

<img width="601" alt="image" src="https://github.com/user-attachments/assets/742d25d8-f3bf-409e-b646-7d01b3f37c3f" />

<img width="601" alt="image" src="https://github.com/user-attachments/assets/916a9c28-8d6d-42d6-a85a-6e9d7fe085cf" />




