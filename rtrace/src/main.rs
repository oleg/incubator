use rtrace::Tuple;

fn main() {
    let e = Environment {
        gravity: Tuple::vector(0.0, -0.1, 0.0),
        wind: Tuple::vector(-0.01, 0.0, 0.0),
    };

    let mut p = Projectile {
        position: Tuple::point(0.0, 1.0, 0.0),
        velocity: Tuple::vector(1.0, 1.0, 0.0).normalize(),
    };

    while p.position.y > 0.0 {
        println!("Projectile position: {:?}", p.position);
        p = tick(e, p);
    }
}

#[derive(Debug, Copy, Clone)]
struct Projectile {
    //point
    position: Tuple,
    //vector
    velocity: Tuple,
}

#[derive(Debug, Copy, Clone)]
struct Environment {
    //vector
    gravity: Tuple,
    //vector
    wind: Tuple,
}

fn tick(env: Environment, proj: Projectile) -> Projectile {
    let position = proj.position + proj.velocity;
    let velocity = proj.velocity + env.gravity + env.wind;
    Projectile { position, velocity }
}