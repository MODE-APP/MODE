//
//  ExplorePageViewController.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/16/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class ExplorePageViewController: UIViewController {
    
    // MARK: - Properties
    
    var loadingPosts: Bool = false
    var loadedAllPosts: Bool = false
    var dataSource: [UIImage] = []
    var numberOfPosts: Int = 300
    
    // MARK: - Outlets
    
    @IBOutlet weak var searchBarPlaceholderLabel: UILabel!
    @IBOutlet weak var searchBarBackgroundColorView: UIView!
    @IBOutlet weak var searchBarTextField: UITextField!
    @IBOutlet weak var catagoryCollectionView: UICollectionView!
    @IBOutlet weak var explorePostsCollectionView: UICollectionView!
    
    // MARK: - Lifecycle
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        for _ in 0..<25 {
            var imageName: String = ""
            let imageNumber = Int.random(in: 0..<5)
            
            switch imageNumber {
            case 0: imageName = "profileOne"
            case 1: imageName = "profileTwo"
            case 2: imageName = "profileThree"
            case 3: imageName = "profileFour"
            case 4: imageName = "profileFive"
            default: fatalError("Big Error at \(#function)")
            }
            let image = UIImage(named: imageName)!
            dataSource.append(image)
        }
        
        searchBarTextField.delegate = self
        catagoryCollectionView.delegate = self
        catagoryCollectionView.dataSource = self
        explorePostsCollectionView.delegate = self
        explorePostsCollectionView.dataSource = self
        updateViews()
    }
    
    // MARK: - Actions
    
    @IBAction func didStartSearching(_ sender: Any) {
        hideSearchBarPlaceholderLabel()
    }
    
    @IBAction func didEndSearching(_ sender: Any) {
        showSearchBarPlaceholderLabel()
    }
    
    // MARK: - Custom Functions
    
    func hideSearchBarPlaceholderLabel() {
        searchBarPlaceholderLabel.isHidden = true
    }
    
    func showSearchBarPlaceholderLabel() {
        searchBarPlaceholderLabel.isHidden = false
    }
    
    func updateViews() {
        searchBarTextField.font = UIFont(name: "SFProDisplay-Light", size: 16)
        searchBarBackgroundColorView.backgroundColor = UIColor(red:0.93, green:0.93, blue:0.93, alpha:1.0)
        searchBarBackgroundColorView.layer.cornerRadius = searchBarTextField.frame.height / 2
    }
    
    func getImages (completion: @escaping () -> Void) {
        let timer = Timer(timeInterval: 0.7, repeats: false) { (_) in
            completion()
        }
        RunLoop.current.add(timer, forMode: .common)
    }
    
    func loadMorePosts() {
        getImages() {
            DispatchQueue.main.async {
                for _ in 0..<30 {
                    var imageName: String = ""
                    let imageNumber = Int.random(in: 0..<5)
                    
                    switch imageNumber {
                    case 0: imageName = "profileOne"
                    case 1: imageName = "profileTwo"
                    case 2: imageName = "profileThree"
                    case 3: imageName = "profileFour"
                    case 4: imageName = "profileFive"
                    default: fatalError("Big Error at \(#function)")
                    }
                    let image = UIImage(named: imageName)!
                    if self.loadedAllPosts == false {
                        self.dataSource.append(image)
                    }
                    if self.dataSource.count == self.numberOfPosts {
                        self.loadedAllPosts = true
                    }
                }
                self.explorePostsCollectionView.reloadData()
                self.loadingPosts = false
            }
        }
    }
    /*
     // MARK: - Navigation
     
     // In a storyboard-based application, you will often want to do a little preparation before navigation
     override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
     // Get the new view controller using segue.destination.
     // Pass the selected object to the new view controller.
     }
     */
    
} // End Of Class

// MARK: - Extensions

extension ExplorePageViewController: UITextFieldDelegate {
    func textFieldShouldReturn(_ textField: UITextField) -> Bool {
        textField.resignFirstResponder()
        textField.text = ""
        return true
    }
}

extension ExplorePageViewController: UICollectionViewDelegate, UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
    
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        if collectionView.restorationIdentifier == "catagoryCollectionView" {
            return 7
        } else {
            return dataSource.count
        }
    }
    
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        if collectionView.restorationIdentifier == "catagoryCollectionView" {
            guard let cell = catagoryCollectionView.dequeueReusableCell(withReuseIdentifier: "catagoryCell", for: indexPath) as? CatagoryCollectionViewCell else {return UICollectionViewCell()}
            return cell
        } else {
            guard let cell = explorePostsCollectionView.dequeueReusableCell(withReuseIdentifier: "photoCell", for: indexPath) as? PostCollectionViewCell else {return UICollectionViewCell()}
            
            cell.image = dataSource[indexPath.row]
            
            return cell
        }
    }
    
    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
        if collectionView.restorationIdentifier == "catagoryCollectionView" {

        } else {

            let size = CGSize(width: explorePostsCollectionView.frame.width / 3 - 1, height: explorePostsCollectionView.frame.width / 3 * 1.6 - 1)
            return size
        }
    }
    
    // MARK: - FlowLayout
    
    //    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
    //        let size = CGSize(width: profileCollectionView.frame.width / 3 - 1, height: profileCollectionView.frame.width / 3 * 1.6 - 1)
    //        return size
    //    }
    
    //    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, referenceSizeForFooterInSection section: Int) -> CGSize {
    //        if loadedAllPosts {
    //            return CGSize.zero
    //        }
    //        return CGSize(width: profileCollectionView.frame.width, height: 55)
    //    }
    
    //    func collectionView(_ collectionView: UICollectionView, viewForSupplementaryElementOfKind kind: String, at indexPath: IndexPath) -> UICollectionReusableView {
    //        guard let footerView = profileCollectionView.dequeueReusableSupplementaryView(ofKind: kind, withReuseIdentifier: "loadDataFooter", for: indexPath) as? PostFooterActivtityIndicatorCollectionReusableView else {fatalError("Couldn't get footer view at \(#function)")}
    //
    //        return footerView
    //    }
    
    //    func scrollViewDidScroll(_ scrollView: UIScrollView) {
    //        let heightFromBottom = scrollView.contentSize.height - scrollView.contentOffset.y
    //        if heightFromBottom < 1300 && loadingPosts == false && loadedAllPosts == false{
    //            loadMorePosts()
    //            loadingPosts = true
    //        }
    //    }
}

