class If < Struct.new(:condition, :consequence, :alternative)
  
  def evaluate environment
    case condition.evaluate environment
    when Boolean.new(true)
      consequence.evaluate environment
    when Boolean.new(false)
      alternative.evaluate environment
    end
  end
  
end
