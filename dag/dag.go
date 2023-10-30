package dag

import queue "github.com/emirpasic/gods/queues/linkedlistqueue"

// DAG 有向无环图
type DAG struct {
	Vertexes []*Vertex
}

// Vertex 顶点
type Vertex struct {
	Key      string
	Value    interface{}
	Parents  []*Vertex
	Children []*Vertex
}

// AddVertex 添加顶点
func (dag *DAG) AddVertex(v *Vertex) {
	dag.Vertexes = append(dag.Vertexes, v)
}

// AddEdge 添加边
func (dag *DAG) AddEdge(from, to *Vertex) {
	from.Children = append(from.Children, to)
	to.Parents = append(to.Parents, from)
}

/*
1. 选择起始节点入队列

2. 节点出队列

2.1 队列空了，说明遍历完毕返回

2.2 已访问跳过，未访问顶点放入输出 slice 中

2.3 将节点的所有未访问邻接节点 Children 放入队列
 3. 重复执行 2
*/
func BFS(root *Vertex) []*Vertex {
	q := queue.New()
	q.Enqueue(root)
	visited := make(map[string]*Vertex)
	all := make([]*Vertex, 0)
	for q.Size() > 0 {
		qSize := q.Size()
		for i := 0; i < qSize; i++ {
			tmp, _ := q.Dequeue()
			currVert := tmp.(*Vertex)
			if _, ok := visited[currVert.Key]; ok {
				continue
			}
			// 记录访问节点
			visited[currVert.Key] = currVert
			all = append([]*Vertex{currVert}, all...)
			for _, val := range currVert.Children {
				if _, ok := visited[val.Key]; !ok {
					q.Enqueue(val) //add child
				}
			}
		}
	}
	return all
}

func DoTasks(vertexes []*Vertex) {

}
