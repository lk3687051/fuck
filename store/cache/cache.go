package cache
import (
	"github.com/patrickmn/go-cache"
)

type Value struct {
	Resource corev2.Resource
	Synth    interface{}
}
