package task_runner

/*
如何写task Runner:
//task_runner这个包是一个具有独立功能的包，可以独自运行的包，而不是纯的方法调用的包
1.首先我们加一个task_runner_main：因为在task_runner中我们既有dispatcher，又有executor，然后我们要将所有的这些任务集中起来。
	我们会再trmain.go中将它们初始化，并将它们跑起来；然后在外面的main中将trmain.go初始化并且启动起来
2.我们需要一个原生的task的包，里面是我们解耦的dispatcher和executor中真正跑的东西
3.我们加一个runner的包，这是比较重要的部分，整个逻辑都在这里面
4.加defs 包
*/
