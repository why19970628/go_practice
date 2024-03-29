工厂方法：该模式有一个抽象基类和若干个派生的具体工厂类，基类定义了一个虚工厂方法返回指定产品类的基类，派生类需要实现该虚方法并创建具体产品类返回。

注意工厂方法的每个具体工厂只负责返回一种产品类。

同样以手机生产做例子：

Nokia工厂方法模式有一个工厂基类NokiaFactory，注意此工厂和上面不一样，是抽象的。

该类定义一个虚工厂方法CreateNokiaPhone，该方法返回NokiaPhone基类。

然后不同型号的手机对应一个该型号的手机工厂，比如N97Factory，注意此工厂派生于NokiaFactory基类，N97Factory实现虚工厂方法，它返回值是具体的Nokia手机类，如new N97Phone。（注意N97Phone是NokiaPhone的派生类）这样的优点就是，新出一个Nokia手机型号，只需派生一个该型号的工厂而无需修改原来的代码。

符合封闭修改，开放扩展原则。




## 工厂方法模式的优点
- 灵活性增强，对于新产品的创建，只需多写一个相应的工厂类。
- 典型的解耦框架。高层模块只需要知道产品的抽象类，无须关心其他实现类，满足迪米特法则、依赖倒置原则和里氏替换原则。


## 工厂方法模式的缺点
- 类的个数容易过多，增加复杂度。
- 增加了系统的抽象性和理解难度。
- 只能生产一种产品，此弊端可使用抽象工厂模式解决。

无论是简单工厂还是工厂方法都只能生产一种产品，如果工厂需要创建生态里的多个产品，就需要更进一步，使用第三级的工厂模式--抽象工厂。