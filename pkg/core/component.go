package core

type Component interface {
    Start() error
    Stop() error
    // 기타 공통 메서드...
}