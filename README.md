## 目录
- [注意](#注意)
- [创建型设计模式](#创建型设计模式)
    - [x] [🏠简单工厂模式 (Simple Factory) ](https://blog.you-tang.com/detail/21)
    - [x] [🏭工厂方法模式 (Factory Method) ](https://blog.you-tang.com/detail/20)
    - [x] [🔨抽象工厂模式 (Abstract Factory) ](https://blog.you-tang.com/detail/19)
    - [x] [👷建造者模式 (Builder Pattern) ](https://blog.you-tang.com/detail/18)
    - [x] [🐑原型模式 (Prototype Pattern) ](https://blog.you-tang.com/detail/17)
    - [x] [💍单例模式 (Singleton Pattern) ](https://blog.you-tang.com/detail/16)
- [结构型设计模式](#结构型设计模式)
    - [ ] [🔌适配器模式（Adapter）](#🔌适配器模式adapter)
    - [ ] [🚡桥梁模式（Bridge）](#🚡桥梁模式bridge)
    - [ ] [🌿组合模式（Composite）](#🌿组合模式composite)
    - [ ] [☕装饰模式（Decorator）](#☕装饰模式decorator)
    - [ ] [📦门面模式（Facade）](#📦门面模式facade)
    - [ ] [🍃享元模式（Flyweight）](#🍃享元模式flyweight)
    - [ ] [🎱代理模式（Proxy）](#🎱代理模式proxy)
- [行为型设计模式](#行为型设计模式)
    - [ ] [🔗责任链模式（Chain Of Responsibilities）](#🔗责任链模式chain-of-responsibilities)
    - [ ] [👮命令行模式（Command）](#👮命令行模式command)
    - [ ] [➿迭代器模式（Iterator）](#➿迭代器模式iterator)
    - [ ] [👽中介者模式（Mediator）](#👽中介者模式mediator)
    - [ ] [💾备忘录模式（Memento）](#💾备忘录模式memento)
    - [ ] [😎观察者模式（Observer）](#😎观察者模式observer)
    - [ ] [🏃访问者模式（Visitor）](#🏃访问者模式visitor)
    - [ ] [💡策略模式（Strategy）](#💡策略模式strategy)
    - [ ] [💢状态模式（State）](#💢状态模式state)
    - [ ] [📒模板方法模式（Template Method）](#📒模板方法模式template-method)
- [总结](#总结)

<!-- /TOC -->

# 介绍
本项目参考与 [design-patterns-for-humans](https://github.com/kamranahmedse/design-patterns-for-humans) 
并将其中的设计模式用 go 语言的方式进行实现。并配备了完整的测试覆盖

## 注意

* 设计模式不是解决所有问题的灵丹妙药。
* 不要试图强迫他们; 如果这样做的话，应该发生坏事。
* 请记住，设计模式是问题的解决方案，而不是解决问题的解决方案；所以不要过分思考。
* 如果以正确的方式在正确的地方使用，他们可以证明是救世主; 否则他们可能会导致代码混乱。