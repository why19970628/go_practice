# 策略模式
策略模式(Strategy Pattern)，一个类的行为或其算法可以在运行时改变。

## 什么是策略模式
策略模式定义一系列算法，把他们一个一个封装起来，并且使他可相互替换。

## 解决了什么问题
解决了在多种算法相似的情况下，使用if...else所带来的复杂和难以维护。
策略模式将这些算法封装成一个一个的类，任意的替换。

## 优点
算法可自由切换
避免多重条件判断
扩展性良好

## 缺点
策略类会增多
所有策略类都需要对外暴露