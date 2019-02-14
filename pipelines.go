package mmio

import "sync"

// many of these pipelines have been taken/modelled from: Concurrency in Go by Katherine Cox-Buday, 2017 (https://katherine.cox-buday.com/concurrency-in-go/)
// code available here: https://github.com/kat-co/concurrency-in-go-src
// errata available here: https://www.oreilly.com/catalog/errata.csp?isbn=0636920046189

// GENERATORS

// IntGenerator is a sample pipeline generator used to return a set of integers into a stream (pg. 106)
func IntGenerator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

// Repeat is a generator that will repeat values sent to it until it's told to stop (pg.109)
func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

// RepeatFn is a generator that repeatedly calls a function (pg. 110)
func RepeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

// Take will only take the first n items off its incoming valueStream (pg. 110)
func Take(done <-chan interface{}, valueStream <-chan interface{}, n int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

// FanOut takes a function and initiates it n times (pg. 116)
// Note: the function must be able to run independently.
// Note: this have yet to be used, and may need testing
func FanOut(done <-chan interface{}, fn func() <-chan interface{}, n int) []chan interface{} {
	fanStream := make([]chan interface{}, n)
	for i := 0; i < n; i++ {
		select {
		case <-done:
			return nil
		case fanStream[i] <- fn():
		}
	}
	return fanStream
}

// FanIn is a means of multiplexing (joining a number of streams into a single stream) (pg. 117)
func FanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexStream <- i:
			}
		}
	}

	// select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexStream)
	}()

	return multiplexStream
}

// OTHER HELPFUL STREAMS

// OrDone channel helps prevent Goroutine leaks when working with channels from disparate parts of code
func OrDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valueStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valueStream
}

// Tee channel take a single channel and returns two separate channels
func Tee(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		for val := range OrDone(done, in) {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}

// Split a channel into n channels that receive messages in a round-robin fashion
// modified from: http://tmrts.com/blog/fan-in--fan-out-messaging-patterns/
func Split(done <-chan interface{}, in <-chan interface{}, n int) []chan interface{} {
	cs := make([]chan interface{}, n)
	for i := 0; i < n; i++ {
		cs[i] = make(chan interface{})
	}

	// Distributes the work in a round robin fashion among the stated number
	// of channels until the main channel has been closed. In that case, close
	// all channels and return.
	distributeToChannels := func(in <-chan interface{}, cs []chan interface{}) {
		// Close every channel when the execution ends.
		defer func(cs []chan interface{}) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for _, c := range cs {
				select {
				case <-done:
					return
				case val, ok := <-in:
					if !ok {
						return
					}
					c <- val
				}
			}
		}

		// for {
		// 	for _, c := range cs {
		// 		select {
		// 		case <-done:
		// 			return
		// 		case c <- <-OrDone(done, in):
		// 		}
		// 	}
		// }
	}

	go distributeToChannels(in, cs)

	return cs
}
