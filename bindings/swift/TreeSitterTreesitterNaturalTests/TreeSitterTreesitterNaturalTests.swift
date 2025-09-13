import XCTest
import SwiftTreeSitter
import TreeSitterTreesitterNatural

final class TreeSitterTreesitterNaturalTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_treesitter_natural())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading TreesitterNatural grammar")
    }
}
