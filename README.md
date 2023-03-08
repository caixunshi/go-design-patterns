# 设计模式golang版本
## 介绍
&emsp;&emsp;在开始使用go语言开发时，因为go没有显示的继承语法，想要实现一些设计模式时总是感到非常别扭，创建这个项目的初衷是希望通过
用go语言实现各种设计模式来提升对go语言的理解，同时也想提供一个文档，让大家在使用go语言实现设计模式有疑问时，通过这个项目能得到一些启发。
<br/>
&emsp;&emsp;希望在github这个平台上跟大家一起交流go语言的使用心得。
## 设计模式简介
&emsp;&emsp;设计模式（Design pattern）代表了最佳的实践，通常被有经验的面向对象的软件开发人员所采用。设计模式是软件开发人员在软件开发过程中
面临的一般问题的解决方案。这些解决方案是众多软件开发人员经过相当长的一段时间的试验和错误总结出来的。
<br/>
## 什么是 GOF（四人帮，全拼 Gang of Four）？
在 1994 年，由 Erich Gamma、Richard Helm、Ralph Johnson 和 John Vlissides 四人合著出版了一本名为 Design Patterns - Elements 
of Reusable Object-Oriented Software（中文译名：设计模式 - 可复用的面向对象软件元素） 的书，该书首次提到了软件开发中设计模式的概念。
四位作者合称 GOF（四人帮，全拼 Gang of Four）。他们所提出的设计模式主要是基于以下的面向对象设计原则。
* 对接口编程而不是对实现编程。
* 优先使用对象组合而不是继承。
<br/>

## 设计模式汇总
根据设计模式的参考书 Design Patterns - Elements of Reusable Object-Oriented Software（中文译名：设计模式 - 可复用的面向对象软件元素） 
中所提到的，总共有 23 种设计模式。这些模式可以分为三大类：创建型模式（Creational Patterns）、结构型模式（Structural Patterns）、
行为型模式（Behavioral Patterns）。

### 创建型模式
创建型设计模式提供了一种在创建对象的同时隐藏创建逻辑的方式，而不是使用 new 运算符直接实例化对象。这使得程序在判断针对某个给定实例需要创建哪些对象时更加灵活。
* <a href="https://github.com/caixunshi/go-design-patterns/tree/main/creational/factory">工厂模式（Factory Pattern）</a>
* <a href="https://github.com/caixunshi/go-design-patterns/tree/main/creational/abstract-factory">抽象工厂模式（Abstract Factory Pattern）</a>
* <a href="https://github.com/caixunshi/go-design-patterns/tree/main/creational/singleton">单例模式（Singleton Pattern）</a>
* <a href="">建造者模式（Builder Pattern）</a>
* <a href="">原型模式（Prototype Pattern）</a>

### 结构型模式
结构型设计模式关注类和对象的组合。继承的概念被用来组合接口和定义组合对象获得新功能的方式。
* <a href="">适配器模式（Adapter Pattern）</a>
* <a href="">桥接模式（Bridge Pattern）</a>
* <a href="">过滤器模式（Filter、Criteria Pattern）</a>
* <a href="">组合模式（Composite Pattern）</a>
* <a href="">装饰器模式（Decorator Pattern）</a>
* <a href="">外观模式（Facade Pattern）</a>
* <a href="">享元模式（Flyweight Pattern）</a>
* <a href="">代理模式（Proxy Pattern）</a>

### 行为型模式
行为型模式不仅仅关注类和对象的结构，而且重点关注它们之间的相互作用。 通过行为型模式，可以更加清晰地划分类与对象的职责，并研究系统在运行时实例对象之间的交互。在系统运行时，对象并不是孤 立的，它们可以通过相互通信与协作完成某些复杂功能，一个对象在运行时也将影响到其他对象的运行。
* <a href="">责任链模式（Chain of Responsibility Pattern）</a>
* <a href="">命令模式（Command Pattern）</a>
* <a href="">解释器模式（Interpreter Pattern）</a>
* <a href="">迭代器模式（Iterator Pattern）</a>
* <a href="">中介者模式（Mediator Pattern）</a>
* <a href="">备忘录模式（Memento Pattern）</a>
* <a href="">观察者模式（Observer Pattern）</a>
* <a href="">状态模式（State Pattern）</a>
* <a href="">空对象模式（Null Object Pattern）</a>
* <a href="">策略模式（Strategy Pattern）</a>
* <a href="">模板模式（Template Pattern）</a>
* <a href="">访问者模式（Visitor Pattern）</a>
