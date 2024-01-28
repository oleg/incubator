package algs.ch10d1;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class StackOnQueuesTest {

    @Test
    void stack_empty_after_creation() {
        var stack = new StackOnQueues<>(5);
        assertTrue(stack.isEmpty());
    }

    @Test
    void stack_is_not_empty_after_push() {
        var stack = new StackOnQueues<>(3);
        stack.push(100);
        assertFalse(stack.isEmpty());
    }

    @Test
    void stack_is_empty_after_push_and_pop() {
        var stack = new StackOnQueues<>(3);
        stack.push(100);
        stack.pop();
        assertTrue(stack.isEmpty());
    }

    @Test
    void cant_pop_empty_stack() {
        var stack = new StackOnQueues<Integer>(10);
        assertThrows(IllegalStateException.class, stack::pop);
    }

    @Test
    void cant_push_full_stack() {
        var stack = new StackOnQueues<Integer>(2);
        stack.push(1);
        stack.push(2);
        assertThrows(IllegalStateException.class, () -> stack.push(3));
    }

    @Test
    void pop_returns_most_recently_pushed_element() {
        var stack = new StackOnQueues<Integer>(10);
        stack.push(100);
        stack.push(200);
        assertEquals(200, stack.pop());
        assertEquals(100, stack.pop());
    }


    @Test
    void stack_is_full() {
        var stack = new StackOnQueues<>(3);
        stack.push(100);
        stack.push(100);
        stack.push(100);

        assertTrue(stack.isFull());
    }

    @Test
    void stack_is_not_full() {
        var stack = new StackOnQueues<>(3);
        stack.push(100);
        stack.push(100);

        assertFalse(stack.isFull());
    }

    @Test
    void complex_test() {
        var stack = new StackOnQueues<>(4);
        stack.push(1);
        stack.push(2);
        stack.push(3);
        stack.push(4);
        assertEquals(4, stack.pop());
        assertEquals(3, stack.pop());
        assertEquals(2, stack.pop());
        assertEquals(1, stack.pop());
        assertTrue(stack.isEmpty());
    }


}